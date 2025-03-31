import { FormEvent } from "react";

type MessageFormProps = {
  input: string;
  setInput: (name: string) => void;
  handleStrikethrough: () => void;
  handleSend: () => void;
  error: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => Promise<void>;
};

export const MessageForm = ({
  input,
  setInput,
  handleStrikethrough,
  handleSend,
  error,
  onSubmit,
}: MessageFormProps) => (
  <form className="border-t p-4 bg-white" onSubmit={onSubmit}>
    <div className="flex gap-2">
      <textarea
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="メッセージを入力..."
        className="flex-1 p-2 border rounded-lg resize-none"
        rows={3}
      />
      <div className="flex flex-col gap-2">
        <button
          type="button"
          onClick={handleStrikethrough}
          className="px-4 py-2 bg-gray-100 rounded hover:bg-gray-200"
        >
          打ち消し線
        </button>
        <button
          type="button"
          onClick={handleSend}
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          送信
        </button>
      </div>
    </div>
  </form>
);
