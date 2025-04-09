"use client";

import { signOut, useSession } from "next-auth/react";
import Link from "next/link";
import { useEffect, useRef, useState } from "react";

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
  const [open, setOpen] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const popoverRef = useRef<HTMLDivElement>(null);
  const modalRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        popoverRef.current &&
        !popoverRef.current.contains(event.target as Node)
      ) {
        setOpen(false);
      }
    };

    const handleModalClickOutside = (event: MouseEvent) => {
      if (
        modalRef.current &&
        !modalRef.current.contains(event.target as Node)
      ) {
        setShowModal(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    document.addEventListener("mousedown", handleModalClickOutside);

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
      document.removeEventListener("mousedown", handleModalClickOutside);
    };
  }, []);

  const handleIntroduceMembers = () => {
    setShowModal(true);
    setOpen(false);
  };

  return (
    <div className="w-1/10 bg-gray-100 p-4 border-r base-color text-white">
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

        <div className="relative inline-block text-left" ref={popoverRef}>
          <button
            onClick={() => setOpen(!open)}
            className="px-4 py-2 rounded hover:bg-gray-600"
          >
            +
          </button>

          {open && (
            <div className="absolute z-10 mt-2 w-48 bg-white border rounded shadow-lg p-4">
              <p>作成</p>
              <p
                className="cursor-pointer hover:bg-gray-100 p-1 rounded"
                onClick={handleIntroduceMembers}
              >
                メンバーを紹介する
              </p>
            </div>
          )}
        </div>
      </div>

      {/* モーダル */}
      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div
            ref={modalRef}
            className="bg-white p-6 rounded-lg shadow-xl w-[600px]"
          >
            <div className="flex justify-between items-center mb-4">
              <h2 className="text-xl font-bold">
                {/* todo:ちゃんとしたものに書き換える */}
                {/* workspacesじゃなく単一のworkspaceをもたせるのでいい気がするので */}
                {workspaces[0].name}にメンバーを招待する
              </h2>
              <button
                onClick={() => setShowModal(false)}
                className="text-gray-500 hover:text-gray-700"
              >
                ✕
              </button>
            </div>
            <textarea
              placeholder="name@gmail.com"
              className="w-full p-2 border rounded resize-none"
            ></textarea>
            <p className="text-center">または</p>
            <button type="button" disabled>
              Google Workspace で続行する
            </button>
            <div className="mt-6 flex justify-between">
              <button>招待リンクをコピーする</button>
              <button
                onClick={() => setShowModal(false)}
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              >
                送信
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
