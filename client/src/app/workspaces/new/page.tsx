"use client";

import { api } from "@/app/lib/api-client";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function NewWorkspace() {
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    try {
      const response = await api.post("/workspaces", {
        email: email,
      });

      if (response.status === 200) {
        router.push("/"); // 作成成功後はホームページにリダイレクト
      }
    } catch (error: any) {
      setError(
        error.response?.data?.error || "ワークスペースの作成に失敗しました"
      );
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-md w-full space-y-8 p-8 bg-white rounded-lg shadow">
        <div>
          <h1 className="text-2xl font-bold text-center">
            ワークスペースを作成
          </h1>
          <p className="mt-2 text-center text-sm text-gray-600">
            メールアドレスを入力してワークスペースを作成してください
          </p>
        </div>
        <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
          <div>
            <label
              htmlFor="email"
              className="block text-sm font-medium text-gray-700"
            >
              メールアドレス
            </label>
            <input
              id="email"
              name="email"
              type="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              placeholder="example@company.com"
            />
          </div>

          {error && <div className="text-red-600 text-sm">{error}</div>}

          <button
            type="submit"
            className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            ワークスペースを作成
          </button>
        </form>
      </div>
    </div>
  );
}
