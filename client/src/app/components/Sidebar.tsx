"use client";

import { signOut, useSession } from "next-auth/react";
import Link from "next/link";

type Workspace = {
  id: number;
  name: string;
  owner_id: number;
  theme: string;
};

interface SidebarProps {
  workspaces: Workspace[];
}

export default function Sidebar({ workspaces }: SidebarProps) {
  const { data: session } = useSession();

  return (
    <div className="w-1/10 bg-gray-100 p-4 border-r">
      <div className="mb-4">
        <Link href="/workspaces/new">ワークスペースを作成</Link>
      </div>

      <div className="mb-4">
        <h2>ワークスペース一覧</h2>
        <ul>
          {workspaces &&
            workspaces.map((workspace) => (
              <li key={workspace.id}>
                <Link
                  href={`/workspaces/${workspace.id}/channels/${workspace.youngestChannelId}`}
                >
                  {workspace.name}
                </Link>
              </li>
            ))}
        </ul>
      </div>

      <div className="mb-4">
        {/* todo;動的な値に対応 */}
        <Link href="/workspaces/1/channels/1">ホーム</Link>
      </div>
      <div className="mb-4">DM</div>
      <div className="mb-4">アクティビティ</div>
      <div className="mb-4">その他</div>

      <div className="flex justify-between items-center mb-4">
        {/* <h1>{session?.user?.name}</h1> */}
        {/* todo: Image使用 */}
        <img src={session?.user?.image} alt="user-image" />
        {/* <button onClick={() => signOut()}>Sign out</button> */}
      </div>
    </div>
  );
}
