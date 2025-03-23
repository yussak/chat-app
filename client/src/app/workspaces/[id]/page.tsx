"use client";

import { api } from "@/app/lib/api-client";
import { useEffect, useState } from "react";
import { useParams } from "next/navigation";

type Workspace = {
  id: number;
  name: string;
  owner_id: number;
  theme: string;
};

export default function Workspace() {
  const [workspace, setWorkspace] = useState<Workspace | null>(null);
  const params = useParams();
  const id = params.id;

  useEffect(() => {
    const fetchWorkspace = async () => {
      try {
        const response = await api.get(`/workspaces/${id}`);
        if (response.status === 200) {
          setWorkspace(response.data);
        }
      } catch (error) {
        console.error("ワークスペースの取得に失敗しました:", error);
      }
    };

    fetchWorkspace();
  }, [id]);

  if (!workspace) {
    return <div>読み込み中...</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">ワークスペース情報</h1>
      <div className="space-y-2">
        <p>ID: {workspace.id}</p>
        <p>名前: {workspace.name}</p>
        <p>オーナーID: {workspace.owner_id}</p>
        <p>テーマ: {workspace.theme}</p>
      </div>
    </div>
  );
}
