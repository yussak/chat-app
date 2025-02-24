"use client";

import { signIn, signOut, useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { api } from "./lib/api-client";
import EmojiPicker from "emoji-picker-react";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";

export default function Home() {
  const { data: session } = useSession();

  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const [showPicker, setShowPicker] = useState(false);

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
    setShowPicker(false);
  };

  const handleStrikethrough = () => {
    const inputElement = document.querySelector(
      'input[type="text"]'
    ) as HTMLInputElement;
    if (!inputElement) return;

    const start = inputElement.selectionStart;
    const end = inputElement.selectionEnd;

    if (start === null || end === null || start === end) return;

    const selectedText = input.slice(start, end);
    const newText =
      input.slice(0, start) + `~~${selectedText}~~` + input.slice(end);

    setInput(newText);
  };

  // console.log(messages);

  // TODO: ## a と入力したらコメント扱いになるのか何も表示されないので対処
  return (
    <div>
      {session ? (
        <>
          <h1>Welcome, {session.user?.name}</h1>
          <button onClick={() => signOut()}>Sign out</button>

          <ul>
            {messages &&
              messages.map((message) => (
                <li key={message.ID} className="border-t-2">
                  {/* {message.ID} */}
                  <p>
                    <img
                      src={message.User.Image}
                      alt="user image"
                      width={100}
                      height={100}
                    />
                    user name:{message.User.Name}
                  </p>
                  <Markdown remarkPlugins={[remarkGfm]}>
                    {message.Content}
                  </Markdown>
                  {/* TODO:任意のリアクションを表示可能にしたい */}
                  <button onClick={() => setShowPicker(!showPicker)}>
                    botan
                  </button>
                  {showPicker && (
                    <div
                      style={{ position: "absolute", top: "40px", zIndex: 10 }}
                    >
                      <EmojiPicker
                        onEmojiClick={(emojiData) => {
                          handleEmojiSelect(message.ID)(emojiData.emoji);
                        }}
                      />
                    </div>
                  )}
                  <div className="flex gap-2">
                    {Object.entries(JSON.parse(message.reactions)).map(
                      ([emoji, count]) => (
                        <button
                          key={emoji}
                          onClick={() => handleAddReaction(message.ID, emoji)}
                        >
                          {emoji} {count}
                        </button>
                      )
                    )}
                    <button onClick={() => setShowPicker(!showPicker)}>
                      + 追加
                    </button>
                  </div>
                  <button
                    onClick={async () => {
                      await api.delete(`/messages/${message.ID}`);
                      setMessages(messages.filter((t) => t.ID !== message.ID));
                    }}
                  >
                    del
                  </button>
                </li>
              ))}
          </ul>

          <div>
            <input
              type="text"
              value={input}
              onChange={(e) => setInput(e.target.value)}
              placeholder="メッセージを入力..."
            />
            <button onClick={handleStrikethrough}>打ち消し線</button>
            <button onClick={handleSend}>送信</button>
          </div>
        </>
      ) : (
        <>
          <h1>Please sign in</h1>
          <button onClick={() => signIn()}>Sign in</button>
        </>
      )}
    </div>
  );
}
