"use client";

import { useParams } from "next/navigation";
import { useState } from "react";
import { useEffect } from "react";
import { api } from "@/app/lib/api-client";
import { useSession } from "next-auth/react";
import EmojiPicker from "emoji-picker-react";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { MessageForm } from "@/app/messages/components/MessageForm";

export default function Channel() {
  const { data: session } = useSession();
  const params = useParams();
  const id = params.id;

  const [channel, setChannel] = useState(null);
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const [activePickerId, setActivePickerId] = useState<number | null>(null);
  const fetchMessages = () => api.get(`/messages?channel_id=${id}`);
  const postMessage = (content: string) =>
    api.post("/messages", {
      content,
      user: {
        id: Number(session?.user?.id),
        name: session?.user?.name,
        image: session?.user?.image,
      },
      channel_id: Number(id),
    });

  useEffect(() => {
    const fetchChannel = async () => {
      const res = await api.get(`/channels/${id}`);
      setChannel(res.data);
    };
    fetchChannel();
    fetchMessages().then((res) => setMessages(res.data));
  }, [id]);

  const handleSend = async () => {
    if (!input.trim()) return;
    const res = await postMessage(input);
    setMessages([...messages, res.data]);
    setInput("");
  };

  const handleStrikethrough = () => {
    const textArea = document.querySelector("textarea") as HTMLTextAreaElement;
    if (!textArea) return;

    const start = textArea.selectionStart;
    const end = textArea.selectionEnd;

    if (start === null || end === null || start === end) return;

    const selectedText = input.slice(start, end);
    const newText =
      input.slice(0, start) + `~~${selectedText}~~` + input.slice(end);

    setInput(newText);
  };

  const handleAddReaction = async (messageId: string, emoji: string) => {
    try {
      await api.post(`/messages/${messageId}/reactions`, {
        user_id: session?.user?.id,
        emoji,
      });
      // リアクション追加後にメッセージ一覧を再取得
      const res = await fetchMessages();
      setMessages(res.data);
    } catch (error) {
      console.error("リアクションの追加に失敗しました", error);
    }
  };

  const handleEmojiSelect = (messageId: string) => (emoji: string) => {
    handleAddReaction(messageId, emoji);
    setActivePickerId(null);
  };

  return (
    <div className="flex h-screen">
      <p>Channel {id}</p>
      <p>name: {channel?.name}</p>
      <div className="w-7/10 flex flex-col">
        {/* メッセージリスト */}
        <div className="flex-1 overflow-y-auto p-4">
          <ul className="space-y-4">
            {messages &&
              messages.map((message) => (
                // todo message card
                <li
                  key={message.id}
                  className="border rounded-lg p-4 bg-white shadow-sm"
                >
                  <div className="flex items-center gap-2 mb-2">
                    <img
                      src={message.user.image}
                      alt="user image"
                      className="w-10 h-10 rounded-full"
                    />
                    <span className="font-semibold">{message.user.name}</span>
                    <span className="text-sm text-gray-500">
                      {message.created_at}
                    </span>
                  </div>
                  <div className="mb-2">
                    <Markdown remarkPlugins={[remarkGfm]}>
                      {message.content}
                    </Markdown>
                  </div>
                  <div className="flex items-center gap-2">
                    {Object.entries(JSON.parse(message.reactions)).map(
                      ([emoji, count]) => (
                        <button
                          key={emoji}
                          onClick={() => handleAddReaction(message.id, emoji)}
                          className="px-2 py-1 bg-gray-100 rounded hover:bg-gray-200"
                        >
                          {emoji} {count}
                        </button>
                      )
                    )}
                    <button
                      onClick={() => setActivePickerId(message.id)}
                      className="px-2 py-1 bg-gray-100 rounded hover:bg-gray-200"
                    >
                      + 追加
                    </button>
                    <button
                      onClick={async () => {
                        await api.delete(`/messages/${message.id}`);
                        setMessages(
                          messages.filter((t) => t.id !== message.id)
                        );
                      }}
                      className="px-2 py-1 bg-red-100 text-red-600 rounded hover:bg-red-200"
                    >
                      delete
                    </button>
                  </div>
                  {activePickerId === message.id && (
                    <div className="absolute z-10">
                      <EmojiPicker
                        onEmojiClick={(emojiData) => {
                          handleEmojiSelect(message.id)(emojiData.emoji);
                        }}
                      />
                    </div>
                  )}
                </li>
              ))}
          </ul>
        </div>

        <MessageForm
          input={input}
          setInput={setInput}
          handleStrikethrough={handleStrikethrough}
          handleSend={handleSend}
          error={""}
          onSubmit={handleSend}
        />
      </div>
    </div>
  );
}
