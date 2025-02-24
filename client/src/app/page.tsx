"use client";

import { signIn, signOut, useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { api } from "./lib/api-client";

export default function Home() {
  const { data: session } = useSession();

  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");

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
  // console.log(messages);
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
                  message content:{message.Content}
                  {/* TODO:任意のリアクションを表示可能にしたい */}
                  <div className="flex gap-2">
                    <button onClick={() => handleAddReaction(message.ID, "👍")}>
                      👍 {JSON.parse(message.reactions)["👍"] || 0}
                    </button>
                    <button onClick={() => handleAddReaction(message.ID, "❤️")}>
                      ❤️ {JSON.parse(message.reactions)["❤️"] || 0}
                    </button>
                    <button onClick={() => handleAddReaction(message.ID, "😄")}>
                      😄 {JSON.parse(message.reactions)["😄"] || 0}
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
