"use client";

import { useParams } from "next/navigation";
import { useState } from "react";
import { useEffect } from "react";
import { api } from "@/app/lib/api-client";
import { useSession } from "next-auth/react";

import { MessageForm } from "@/app/messages/components/MessageForm";
import { MessageItem } from "@/app/messages/components/MessageItem";
import ChannelList from "@/app/components/ChannelList";

interface Message {
  id: number;
  content: string;
  created_at: string;
  channel_id: number;
  user_id: number;
  updated_at: string;
  user: {
    id: number;
    name: string;
    image: string;
  };
}

interface Channel {
  id: number;
  name: string;
  created_at: string;
}

interface Workspace {
  id: number;
  name: string;
  channels: {
    id: number;
    workspace_id: number;
    name: string;
    is_public: boolean;
  }[];
}

export default function Channel() {
  const { data: session } = useSession();
  const params = useParams();
  const id = params.channelId;
  const workspaceId = params.workspaceId;

  const [workspace, setWorkspace] = useState<Workspace | null>(null);
  const [channel, setChannel] = useState<Channel | null>(null);
  const [messages, setMessages] = useState<Message[]>([]);
  const [message, setMessage] = useState("");
  const [activePickerId, setActivePickerId] = useState<number | null>(null);
  const fetchMessages = () => api.get(`/messages?channel_id=${id}`);
  const postMessage = (content: string) =>
    api.post("/messages", {
      content,
      user: {
        id: Number(session?.user?.id),
        name: session?.user?.name,
        image: session?.user?.image,
      },
      channel_id: Number(id),
    });

  useEffect(() => {
    const fetchWorkspace = async () => {
      const res = await api.get(`/workspaces/${workspaceId}`);
      setWorkspace(res.data);
    };
    fetchWorkspace();

    const fetchChannel = async () => {
      const res = await api.get(`/channels/${id}`);
      setChannel(res.data);
    };
    fetchChannel();
    fetchMessages().then((res) => setMessages(res.data));
  }, [id, workspaceId]);

  const handleSend = async () => {
    if (!message.trim()) return;
    const res = await postMessage(message);

    if (messages == null) {
      setMessages([res.data]);
      setMessage("");
      return;
    }
    setMessages([...messages, res.data]);
    setMessage("");
  };

  const handleStrikethrough = () => {
    const textArea = document.querySelector("textarea") as HTMLTextAreaElement;
    if (!textArea) return;

    const start = textArea.selectionStart;
    const end = textArea.selectionEnd;

    if (start === null || end === null || start === end) return;

    const selectedText = message.slice(start, end);
    const newText =
      message.slice(0, start) + `~~${selectedText}~~` + message.slice(end);

    setMessage(newText);
  };

  const handleAddReaction = async (messageId: number, emoji: string) => {
    try {
      await api.post(`/messages/${messageId}/reactions`, {
        user_id: session?.user?.id,
        emoji,
      });
      // リアクション追加後にメッセージ一覧を再取得
      const res = await fetchMessages();
      setMessages(res.data);
    } catch (error) {
      console.error("リアクションの追加に失敗しました", error);
    }
  };

  const handleEmojiSelect = (messageId: number) => (emoji: string) => {
    handleAddReaction(messageId, emoji);
    setActivePickerId(null);
  };

  const handleDelete = async (messageId: number) => {
    await api.delete(`/messages/${messageId}`);
    setMessages(messages.filter((message) => message.id !== messageId));
  };

  return (
    <div className="flex h-screen">
      <div className="w-2/10 bg-gray-100 p-4 border-r">
        <h2>{workspace?.name}</h2>
        <p>チャンネル</p>
        {workspace && (
          <ChannelList
            channels={workspace.channels}
            workspaceId={workspace.id}
          />
        )}
      </div>
      <div className="w-full flex flex-col">
        <div className="p-4 border-b">
          <h1 className="text-xl font-bold">#{channel?.name}</h1>
        </div>
        {/* メッセージリスト */}
        <div className="flex-1 overflow-y-auto p-4">
          <ul className="space-y-4">
            {messages &&
              messages.map((message) => (
                <MessageItem
                  key={message.id}
                  message={message}
                  handleAddReaction={handleAddReaction}
                  activePickerId={activePickerId}
                  setActivePickerId={setActivePickerId}
                  handleDelete={handleDelete}
                  handleEmojiSelect={handleEmojiSelect}
                />
              ))}
          </ul>
        </div>

        <MessageForm
          message={message}
          placeholder={`#${channel?.name} へのメッセージ`}
          setMessage={setMessage}
          handleStrikethrough={handleStrikethrough}
          handleSend={handleSend}
          error={""}
          onSubmit={handleSend}
        />
      </div>
    </div>
  );
}
