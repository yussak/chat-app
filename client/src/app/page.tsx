"use client";

import axios from "axios";
import { signIn, signOut, useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { api } from "./lib/api-client";

export default function Home() {
  const { data: session } = useSession();

  const [todos, setTodos] = useState([]);
  const [input, setInput] = useState("");

  const fetchtodos = () => api.get("http://localhost:8080/todos");
  const postMessage = (name: string) =>
    api.post("http://localhost:8080/todos", {
      name,
      user: {
        id: session?.user?.id,
        name: session?.user?.name,
        image: session?.user?.image,
      },
    });

  useEffect(() => {
    fetchtodos().then((res) => setTodos(res.data));
  }, []);

  const handleSend = async () => {
    if (!input.trim()) return;
    const res = await postMessage(input);
    setTodos([...todos, res.data]);
    setInput("");
  };

  return (
    <div>
      {session ? (
        <>
          <h1>Welcome, {session.user?.name}</h1>
          <button onClick={() => signOut()}>Sign out</button>

          <ul>
            {todos &&
              todos.map((todo) => (
                <li key={todo.ID} className="border-t-2">
                  {/* {todo.ID} */}
                  <p>
                    <img
                      src={todo.User.Image}
                      alt="user image"
                      width={100}
                      height={100}
                    />
                    user name:{todo.User.Name}
                  </p>
                  todo name:{todo.Name}
                  <button
                    onClick={async () => {
                      await api.delete(
                        `http://localhost:8080/todos/${todo.ID}`
                      );
                      setTodos(todos.filter((t) => t.ID !== todo.ID));
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
