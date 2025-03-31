"use client";

import { signIn, useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { api } from "./lib/api-client";
import Sidebar from "./components/Sidebar";

export default function Home() {
  const { data: session } = useSession();
  const [workspaces, setWorkspaces] = useState([]);
  const fetchWorkspaces = () => api.get("/workspaces");

  useEffect(() => {
    fetchWorkspaces().then((res) => setWorkspaces(res.data));
  }, []);

  return (
    <div className="min-h-screen">
      {session ? (
        <div className="flex h-screen">
          <Sidebar workspaces={workspaces} />
        </div>
      ) : (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
          <div className="text-center">
            <h1 className="text-2xl font-bold mb-4">Please sign in</h1>
            <button
              onClick={() => signIn()}
              className="px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
            >
              Sign in
            </button>
          </div>
        </div>
      )}
    </div>
  );
}
