"use client";

import { api } from "@/app/lib/api-client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { EmailForm } from "../components/EmailForm";
import { NameForm } from "../components/NameForm";
import { DisplayNameForm } from "../components/DisplayNameForm";
import { InvitationForm } from "../components/InvitationForm";
import { ThemeForm } from "../components/ThemeForm";
import { CreateForm } from "../components/CreateForm";

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

    // if (!invitation.trim()) {
    //   setError("招待コードを入力してください");
    //   return;
    // }

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
        theme,
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
    <NameForm
      name={name}
      setName={setName}
      error={error}
      onSubmit={handleNameSubmit}
    />
  );

  const renderDisplayNameForm = () => (
    <DisplayNameForm
      displayName={displayName}
      setDisplayName={setDisplayName}
      error={error}
      onSubmit={handleDisplayNameSubmit}
    />
  );

  const renderInvitationForm = () => (
    <InvitationForm error={error} onSubmit={handleInvitationSubmit} />
  );

  const renderThemeForm = () => (
    <ThemeForm
      theme={theme}
      setTheme={setTheme}
      error={error}
      onSubmit={handleThemeSubmit}
    />
  );

  const renderCreateForm = () => (
    <CreateForm error={error} onSubmit={handleStartSubmit} />
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
        return renderCreateForm();
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
