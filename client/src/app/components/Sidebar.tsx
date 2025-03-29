"use client";

import { signOut, useSession } from "next-auth/react";
import Link from "next/link";

export default function Sidebar({ workspaces }) {
  const { data: session } = useSession();

  return (
    <div className="w-3/10 bg-gray-100 p-4 border-r">
      <div className="flex justify-between items-center mb-4">
        <h1>Welcome, {session?.user?.name}</h1>
        <button onClick={() => signOut()}>Sign out</button>
      </div>

      <div className="mb-4">
        <Link href="/workspaces/new">ワークスペースを作成</Link>
      </div>

      <div className="mb-4">
        <h2>ワークスペース一覧</h2>
        <ul>
          {workspaces &&
            workspaces.map((workspace) => (
              <li key={workspace.id}>
                <Link href={`/workspaces/${workspace.id}`}>
                  {workspace.name}
                </Link>
              </li>
            ))}
        </ul>
      </div>
    </div>
  );
}
