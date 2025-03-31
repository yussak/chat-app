"use client";

import { api } from "@/app/lib/api-client";
import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import Link from "next/link";

type Workspace = {
  id: number;
  name: string;
  owner_id: number;
  theme: string;
  channels: {
    id: number;
    name: string;
  }[];
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
    <div className="flex h-screen">
      {/* サイドバー */}
      <ul className="w-1/10 p-4 border-r">
        <li>チーム名</li>
        <li>ホーム</li>
        <li>DM</li>
        <li>アクティビティ</li>
        <li>その他</li>
      </ul>
      <div className="w-2/10 bg-gray-100 p-4 border-r">
        <h2>{workspace.name}</h2>
        <p>チャンネル</p>
        <ul>
          {workspace.channels &&
            workspace.channels.map((channel) => (
              <li key={channel.id}>
                <Link
                  href={`/channels/${channel.id}`}
                  className="block hover:bg-gray-200 p-2 rounded"
                >
                  # {channel.name}
                </Link>
              </li>
            ))}
        </ul>
      </div>

      {/* メインコンテンツ */}
      <div className="w-7/10 p-4">
        <h1 className="text-2xl font-bold mb-4">ワークスペース情報</h1>
        <div className="space-y-2">
          <p>ID: {workspace.id}</p>
          <p>名前: {workspace.name}</p>
          <p>オーナーID: {workspace.owner_id}</p>
          <p>テーマ: {workspace.theme}</p>
        </div>
      </div>
    </div>
  );
}
