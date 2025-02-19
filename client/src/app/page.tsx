"use client";

import axios from "axios";
import { signIn, signOut, useSession } from "next-auth/react";
import { useEffect, useState } from "react";

export default function Home() {
  const { data: session } = useSession();

  const [todos, setTodos] = useState([]);
  const [input, setInput] = useState("");

  const fetchtodos = () => axios.get("http://localhost:8080/todos");
  const postMessage = (name: string) =>
    axios.post("http://localhost:8080/todos", { name });

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
            {todos.map((todo) => (
              <li key={todo.ID}>
                {todo.ID}
                {todo.Name}
                <button
                  onClick={async () => {
                    await axios.delete(
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
