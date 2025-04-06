import React from "react";
import Header from "@/app/components/Header";
import Sidebar from "@/app/components/Sidebar";

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
    <div style={{ display: "flex", flexDirection: "column", height: "100vh" }}>
      <Header />
      <div style={{ display: "flex", flex: 1 }}>
        <Sidebar workspaceId={workspaceId} />
        <main style={{ flex: 1 }}>{children}</main>
      </div>
    </div>
  );
}
