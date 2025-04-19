import EmojiPicker from "emoji-picker-react";
import Markdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { Message } from "@/app/types";
import { useSession } from "next-auth/react";

type MessageItemProps = {
  message: Message;
  handleAddReaction: (messageId: number, emoji: string) => void;
  activePickerId: number | null;
  setActivePickerId: (id: number) => void;
  handleDelete: (messageId: number) => void;
  handleEmojiSelect: (messageId: number) => (emoji: string) => void;
};

export const MessageItem = ({
  message,
  handleAddReaction,
  activePickerId,
  setActivePickerId,
  handleDelete,
  handleEmojiSelect,
}: MessageItemProps) => {
  const { data: session } = useSession();

  return (
    <li className="border rounded-lg p-4 bg-white shadow-sm">
      <div className="flex items-center gap-2 mb-2">
        <img
          src={message.user.image}
          alt="user image"
          className="w-10 h-10 rounded-full"
        />
        <span className="font-semibold">{message.user.name}</span>
        <span className="text-sm text-gray-500">{message.created_at}</span>
      </div>
      <div className="mb-2">
        <Markdown remarkPlugins={[remarkGfm]}>{message.content}</Markdown>
      </div>
      <div className="flex items-center gap-2">
        {Object.entries(JSON.parse(message.reactions)).map(([emoji, count]) => (
          <button
            key={emoji}
            onClick={() => handleAddReaction(message.id, emoji)}
            className="px-2 py-1 bg-gray-100 rounded hover:bg-gray-200"
          >
            {emoji} {count}
          </button>
        ))}
        <button
          onClick={() => setActivePickerId(message.id)}
          className="px-2 py-1 bg-gray-100 rounded hover:bg-gray-200"
        >
          + 追加
        </button>
        {message.user.id === session?.user?.id && (
          <button
            type="button"
            onClick={() => handleDelete(message.id)}
            className="px-2 py-1 bg-red-100 text-red-600 rounded hover:bg-red-200"
          >
            delete
          </button>
        )}
      </div>
      {activePickerId === message.id && (
        <div className="absolute z-10">
          <EmojiPicker
            onEmojiClick={(emojiData) => {
              handleEmojiSelect(message.id)(emojiData.emoji);
            }}
          />
        </div>
      )}
    </li>
  );
};
