"use client";

import { api } from "@/app/lib/api-client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { EmailForm } from "../components/EmailForm";

type Step = "email" | "name" | "displayName" | "invitation" | "theme" | "start";

// TODO:各ステップにバリデーション追加
// TODO: 表示名を入力させるのはownerだからかというよりworkspaceに参加するときにやるべきな気がするので確認
export default function NewWorkspace() {
  const [step, setStep] = useState<Step>("email");
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [invitation, setInvitation] = useState("");
  const [displayName, setDisplayName] = useState("");
  const [theme, setTheme] = useState("");
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

    setStep("invitation");
  };

  const handleInvitationSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!invitation.trim()) {
      setError("招待コードを入力してください");
      return;
    }

    // TODO:招待できるように実装

    setStep("theme");
  };

  const handleThemeSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!theme.trim()) {
      setError("テーマを入力してください");
      return;
    }

    setStep("start");
  };

  const handleStartSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

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
    <EmailForm
      email={email}
      setEmail={setEmail}
      error={error}
      onSubmit={handleEmailSubmit}
    />
  );

  const renderNameForm = () => (
    <form className="mt-8 space-y-6" onSubmit={handleNameSubmit}>
      <p>手順 1/5</p>
      <div>
        <input
          id="name"
          name="name"
          type="text"
          required
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          placeholder="例: ABC 営業部、ABC 社"
        />
      </div>

      {error && <div className="text-red-600 text-sm">{error}</div>}

      <div className="flex gap-4">
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
      <p>手順 2/5</p>
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
          type="submit"
          className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          次へ
        </button>
      </div>
    </form>
  );

  const renderInvitationForm = () => (
    <form className="mt-8 space-y-6" onSubmit={handleInvitationSubmit}>
      <p>手順 3/5</p>
      <div>
        <label
          htmlFor="invitation"
          className="block text-sm font-medium text-gray-700"
        >
          一緒に仕事をする人をメールアドレスで追加する
        </label>
        <textarea
          id="invitation"
          name="invitation"
          rows={3}
          value={invitation}
          onChange={(e) => setInvitation(e.target.value)}
          className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          placeholder="例: ellis@gmail.com、 maria@gmail.com"
        />
      </div>

      {error && <div className="text-red-600 text-sm">{error}</div>}

      <div className="flex gap-4">
        <button
          // type="submit"
          className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          次へ（未実装）
        </button>
        <button
          type="button"
          onClick={() => setStep("theme")}
          className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          この手順をスキップする
        </button>
      </div>
    </form>
  );

  const renderThemeForm = () => (
    <form className="mt-8 space-y-6" onSubmit={handleThemeSubmit}>
      <p>手順 4/5</p>
      <div>
        <p className="block text-sm font-medium text-gray-700">
          チームで取り組んでいることを教えてください。プロジェクト、キャンペーン、イベント、まとめようとしている案件など、いろいろなことが考えられます。
        </p>
        <input
          id="theme"
          name="theme"
          type="text"
          required
          value={theme}
          placeholder="例: 第４半期予算、秋のキャンペーン"
          onChange={(e) => setTheme(e.target.value)}
          className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        />
        {error && <div className="text-red-600 text-sm">{error}</div>}

        <button
          type="button"
          onClick={() => setStep("start")}
          className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
        >
          次へ
        </button>
      </div>
    </form>
  );

  const renderStartForm = () => (
    <form className="mt-8 space-y-6" onSubmit={handleStartSubmit}>
      <p>ワークスペースの準備ができました！✨</p>
      <button
        type="submit"
        className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        ワークスペースを開始する
      </button>
    </form>
  );

  const getStepTitle = () => {
    switch (step) {
      case "email":
        return "メールアドレスを入力してください";
      case "name":
        return "ワークスペース名を入力してください";
      case "displayName":
        return "表示名を入力してください";
      case "invitation":
        return `${name} チームにはほかに誰がいますか？`;
      case "theme":
        return `チームで今取り組んでいることは何ですか？`;
      case "start":
        return "ワークスペースを開始する";
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
      case "invitation":
        return renderInvitationForm();
      case "theme":
        return renderThemeForm();
      case "start":
        return renderStartForm();
      default:
        return null;
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="max-w-md w-full space-y-8 p-8 bg-white rounded-lg shadow">
        <div>
          <h1 className="text-2xl font-bold text-center">{getStepTitle()}</h1>
        </div>
        {renderForm()}
      </div>
    </div>
  );
}
