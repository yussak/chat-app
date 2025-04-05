// app/workspace/[workspaceId]/layout.tsx
import Sidebar from "@/app/components/Sidebar";
import React from "react";

export default async function WorkspaceLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { workspaceId: string };
}) {
  // paramsが用意されるのを待つ
  const { workspaceId } = await params;
  return (
    <div style={{ display: "flex", height: "100vh" }}>
      <Sidebar workspaceId={workspaceId} />
      <main style={{ flex: 1 }}>{children}</main>
    </div>
  );
}
