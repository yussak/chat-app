import { FormEvent } from "react";

type MessageFormProps = {
  message: string;
  placeholder: string;
  setMessage: (name: string) => void;
  handleStrikethrough: () => void;
  handleSend: () => void;
  error: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => Promise<void>;
};

export const MessageForm = ({
  message,
  placeholder,
  setMessage,
  handleStrikethrough,
  handleSend,
  error,
  onSubmit,
}: MessageFormProps) => (
  <form
    className="flex flex-col border rounded-lg m-4 p-4 bg-white"
    onSubmit={onSubmit}
  >
    {/* 上段ボタン */}
    <div className="flex gap-2 mb-2">
      {/* todo:太字実装 */}
      <button type="button" className="hover:bg-gray-200 font-bold">
        B
      </button>
      {/* todo:実装 */}
      <button type="button" className="px-2 hover:bg-gray-200 italic">
        I
      </button>
      <button
        type="button"
        onClick={handleStrikethrough}
        className="px-2 rounded hover:bg-gray-200"
      >
        {/* 打ち消し線 */}
        <span className="relative font-bold text-sm">
          S
          <span className="absolute left-0 right-0 top-1/2 border-t border-black rotate-[15deg]"></span>
        </span>

        <span className="px-4">|</span>
      </button>
    </div>
    {/* テキストエリア */}
    <textarea
      value={message}
      onChange={(e) => setMessage(e.target.value)}
      placeholder={placeholder}
      className="flex-1 p-2 rounded resize-none mb-2"
      rows={3}
    />
    {/* 下段ボタン */}
    <div className="flex gap-2">
      <div className="flex flex-col gap-2">
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
