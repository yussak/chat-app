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
  const [openPopover, setOpenPopover] = useState(false);
  const [addChannelModal, setAddChannelModal] = useState(false);
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

  const handleAddChannel = () => {
    setAddChannelModal(true);
    setOpenPopover(false);
  };

  return (
    <div className="flex h-screen">
      <div className="w-2/10 bg-gray-100 p-4 border-r aa text-white">
        <h2>{workspace?.name}</h2>
        <p>チャンネル</p>
        {workspace && (
          <>
            <ChannelList
              channels={workspace.channels}
              workspaceId={workspace.id}
            />
            <div className="relative">
              <button onClick={() => setOpenPopover(!openPopover)}>
                チャンネルを追加する
              </button>
            </div>

            {openPopover && (
              <div className="absolute bg-white p-4 rounded shadow-lg">
                <p>
                  <button onClick={handleAddChannel}>
                    新しいチャンネルを作成する
                  </button>
                </p>
                <p>
                  <button>チャンネル一覧</button>
                </p>
              </div>
            )}

            {addChannelModal && (
              <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
                <div className="bg-white p-6 rounded-lg shadow-xl w-[500px]">
                  <h2 className="text-xl font-bold mb-4">
                    チャンネルを追加する
                  </h2>
                  <div className="mb-4">
                    <input
                      type="text"
                      placeholder="チャンネル名"
                      className="w-full p-2 border rounded"
                    />
                  </div>
                  <div className="flex justify-end">
                    <button
                      onClick={() => setAddChannelModal(false)}
                      className="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300 mr-2"
                    >
                      キャンセル
                    </button>
                    <button className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">
                      作成
                    </button>
                  </div>
                </div>
              </div>
            )}
          </>
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
