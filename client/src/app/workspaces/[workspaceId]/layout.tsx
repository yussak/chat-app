import React from "react";
import Header from "@/app/components/Header";
import Sidebar from "@/app/components/Sidebar";
import axios from "axios";
export default async function WorkspaceLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { workspaceId: string };
}) {
  // serverと指定する必要があるので一旦直がきしている
  const response = await axios.get(`http://server:8080/workspaces`);
  const workspaces = response.data;

  return (
    <div style={{ display: "flex", flexDirection: "column", height: "100vh" }}>
      <Header />
      <div style={{ display: "flex", flex: 1 }}>
        <Sidebar workspaces={workspaces} />
        <main style={{ flex: 1 }}>{children}</main>
      </div>
    </div>
  );
}
