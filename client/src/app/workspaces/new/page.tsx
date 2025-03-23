"use client";

import { api } from "@/app/lib/api-client";
import { useState } from "react";
import { useRouter } from "next/navigation";

type Step = "email" | "name" | "displayName";

// TODO:各ステップにバリデーション追加
// TODO: 表示名を入力させるのはownerだからかというよりworkspaceに参加するときにやるべきな気がするので確認
export default function NewWorkspace() {
  const [step, setStep] = useState<Step>("email");
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [displayName, setDisplayName] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  const handleEmailSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    try {
      const response = await api.get(`/users/exists?email=${email}`);
      if (response.status === 200) {
        setStep("name");
      }
    } catch (error: any) {
      setError(error.response?.data?.error || "ユーザーが見つかりませんでした");
    }
  };

  const handleNameSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!name.trim()) {
      setError("ワークスペース名を入力してください");
      return;
    }

    setStep("displayName");
  };

  // 表示名の設定はownerだからかというよりworkspaceに参加するときにやるべきな気がすると思ってるが、ワークスペース作成時には常に初めて参加なのでこれはこれで良さそう
  const handleDisplayNameSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!displayName.trim()) {
      setError("表示名を入力してください");
      return;
    }

    try {
      const response = await api.post("/workspaces", {
        email,
        name,
        displayName,
      });

      if (response.status === 200) {
        router.push("/");
      }
    } catch (error: any) {
      setError(
        error.response?.data?.error || "ワークスペースの作成に失敗しました"
      );
    }
  };

  const renderEmailForm = () => (
    <form className="mt-8 space-y-6" onSubmit={handleEmailSubmit}>
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
        次へ
      </button>
    </form>
  );

  const renderNameForm = () => (
    <form className="mt-8 space-y-6" onSubmit={handleNameSubmit}>
      <div>
        <label
          htmlFor="name"
          className="block text-sm font-medium text-gray-700"
        >
          ワークスペース名
        </label>
        <input
          id="name"
          name="name"
          type="text"
          required
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          placeholder="例: プロジェクト名"
        />
      </div>

      {error && <div className="text-red-600 text-sm">{error}</div>}

      <div className="flex gap-4">
        <button
          type="button"
          onClick={() => setStep("email")}
          className="flex-1 py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          戻る
        </button>
        <button
          type="submit"
          className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          次へ
        </button>
      </div>
    </form>
  );

  const renderDisplayNameForm = () => (
    <form className="mt-8 space-y-6" onSubmit={handleDisplayNameSubmit}>
      <div>
        <label
          htmlFor="displayName"
          className="block text-sm font-medium text-gray-700"
        >
          表示名
        </label>
        <input
          id="displayName"
          name="displayName"
          type="text"
          required
          value={displayName}
          onChange={(e) => setDisplayName(e.target.value)}
          className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          placeholder="例: 山田 太郎"
        />
      </div>

      {error && <div className="text-red-600 text-sm">{error}</div>}

      <div className="flex gap-4">
        <button
          type="button"
          onClick={() => setStep("name")}
          className="flex-1 py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          戻る
        </button>
        <button
          type="submit"
          className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          作成
        </button>
      </div>
    </form>
  );

  const getStepTitle = () => {
    switch (step) {
      case "email":
        return "メールアドレスを入力してワークスペースを作成してください";
      case "name":
        return "ワークスペースの名前を入力してください";
      case "displayName":
        return "ワークスペースでの表示名を設定してください";
      default:
        return "";
    }
  };

  const renderForm = () => {
    switch (step) {
      case "email":
        return renderEmailForm();
      case "name":
        return renderNameForm();
      case "displayName":
        return renderDisplayNameForm();
      default:
        return null;
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
            {getStepTitle()}
          </p>
        </div>
        {renderForm()}
      </div>
    </div>
  );
}
