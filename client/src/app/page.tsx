"use client";

import { signIn, signOut, useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { api } from "./lib/api-client";
import EmojiPicker from "emoji-picker-react";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import Link from "next/link";

export default function Home() {
  const { data: session } = useSession();

  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const [activePickerId, setActivePickerId] = useState<number | null>(null);

  const fetchMessages = () => api.get("/messages");
  const postMessage = (content: string) =>
    api.post("/messages", {
      content,
      user: {
        id: session?.user?.id,
        name: session?.user?.name,
        image: session?.user?.image,
      },
    });

  useEffect(() => {
    fetchMessages().then((res) => setMessages(res.data));
  }, []);

  const handleSend = async () => {
    if (!input.trim()) return;
    const res = await postMessage(input);
    setMessages([...messages, res.data]);
    setInput("");
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

  return (
    <div className="min-h-screen">
      {session ? (
        <div className="flex h-screen">
          {/* サイドバー */}
          <div className="w-3/10 bg-gray-100 p-4 border-r">
            <div className="flex justify-between items-center mb-4">
              <h1>Welcome, {session.user?.name}</h1>
              <button onClick={() => signOut()}>Sign out</button>
            </div>

            <div className="mb-4">
              <Link href="/workspaces/new">ワークスペースを作成</Link>
            </div>
          </div>

          {/* メインコンテンツ */}
          <div className="w-7/10 flex flex-col">
            {/* メッセージリスト */}
            <div className="flex-1 overflow-y-auto p-4">
              <ul className="space-y-4">
                {messages &&
                  messages.map((message) => (
                    <li
                      key={message.ID}
                      className="border rounded-lg p-4 bg-white shadow-sm"
                    >
                      <div className="flex items-center gap-2 mb-2">
                        <img
                          src={message.User.Image}
                          alt="user image"
                          className="w-10 h-10 rounded-full"
                        />
                        <span className="font-semibold">
                          {message.User.Name}
                        </span>
                        <span className="text-sm text-gray-500">
                          {message.CreatedAt}
                        </span>
                      </div>
                      <div className="mb-2">
                        <Markdown remarkPlugins={[remarkGfm]}>
                          {message.Content}
                        </Markdown>
                      </div>
                      <div className="flex items-center gap-2">
                        {Object.entries(JSON.parse(message.reactions)).map(
                          ([emoji, count]) => (
                            <button
                              key={emoji}
                              onClick={() =>
                                handleAddReaction(message.ID, emoji)
                              }
                              className="px-2 py-1 bg-gray-100 rounded hover:bg-gray-200"
                            >
                              {emoji} {count}
                            </button>
                          )
                        )}
                        <button
                          onClick={() => setActivePickerId(message.ID)}
                          className="px-2 py-1 bg-gray-100 rounded hover:bg-gray-200"
                        >
                          + 追加
                        </button>
                        <button
                          onClick={async () => {
                            await api.delete(`/messages/${message.ID}`);
                            setMessages(
                              messages.filter((t) => t.ID !== message.ID)
                            );
                          }}
                          className="px-2 py-1 bg-red-100 text-red-600 rounded hover:bg-red-200"
                        >
                          delete
                        </button>
                      </div>
                      {activePickerId === message.ID && (
                        <div className="absolute z-10">
                          <EmojiPicker
                            onEmojiClick={(emojiData) => {
                              handleEmojiSelect(message.ID)(emojiData.emoji);
                            }}
                          />
                        </div>
                      )}
                    </li>
                  ))}
              </ul>
            </div>

            {/* メッセージ入力エリア */}
            <div className="border-t p-4 bg-white">
              <div className="flex gap-2">
                <textarea
                  value={input}
                  onChange={(e) => setInput(e.target.value)}
                  placeholder="メッセージを入力..."
                  className="flex-1 p-2 border rounded-lg resize-none"
                  rows={3}
                />
                <div className="flex flex-col gap-2">
                  <button
                    onClick={handleStrikethrough}
                    className="px-4 py-2 bg-gray-100 rounded hover:bg-gray-200"
                  >
                    打ち消し線
                  </button>
                  <button
                    onClick={handleSend}
                    className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                  >
                    送信
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      ) : (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
          <div className="text-center">
            <h1 className="text-2xl font-bold mb-4">Please sign in</h1>
            <button
              onClick={() => signIn()}
              className="px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
            >
              Sign in
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
