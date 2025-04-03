// app/workspace/[workspaceId]/layout.tsx
import Sidebar from "@/app/components/Sidebar";
import React from "react";

export default function WorkspaceLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { workspaceId: string };
}) {
  return (
    <div style={{ display: "flex", height: "100vh" }}>
      <Sidebar workspaceId={params.workspaceId} />
      <main style={{ flex: 1 }}>{children}</main>
    </div>
  );
}
