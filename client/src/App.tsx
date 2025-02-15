import { useEffect, useState } from 'react'
import { Todo } from './types';

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [todo, setTodoName] = useState<string | null>(null);

  useEffect(() => {
    fetch('http://localhost:8080/')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then((data) => setTodos(data))
      .catch((error) => console.error('Error fetching data:', error));
  }, []);

  const handleAddTodo = async () => {
    try {
      const response = await fetch("http://localhost:8080/add", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ Name: todo }),
      });

      if (response.ok) {
        const newTodo = await response.json();
        setTodos((prevTodos) => [...prevTodos, newTodo]);
        setTodoName(null);
      } else {
        const errorText = await response.text();
        console.error(`エラー: ${errorText}`);
      }
    } catch (error) {
      console.error("通信エラー:", error);
    }
  };

  const handleDeleteTodo = async (id: number) => {
    try {
      const response = await fetch(`http://localhost:8080/delete?id=${id}`, {
        method: "DELETE",
      })

      if (response.ok) {
        setTodos((prevTodos) => prevTodos.filter((todo) => todo.ID !== id))
      } else {
        const errorText = await response.text()
        console.error(`エラー: ${errorText}`);
      }
    } catch (error) {
      console.error("通信エラー:", error)
    }
  }

  return (
    <>
      <h1>TODOリスト</h1>
      <input
        type="text"
        name="todo"
        onChange={(e) => setTodoName(e.target.value)} // onChangeで入力値を更新
        placeholder="タスクを入力"
      />
      <button onClick={handleAddTodo}>追加</button>
      <ul>
        {todos && todos.length > 0 ? (
          todos.map((todo) => (
            <li key={todo.ID}>
              {todo.Name}<button onClick={() => handleDeleteTodo(todo.ID)}>削除</button>
            </li>
          ))
        ) : (
          <p>Todoがありません</p>
        )}
      </ul>
    </>
  )
}

export default App
