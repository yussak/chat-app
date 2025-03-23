import { FormEvent } from "react";

type CreateFormProps = {
  error: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => Promise<void>;
};

export const CreateForm = ({ error, onSubmit }: CreateFormProps) => (
  <form className="mt-8 space-y-6" onSubmit={onSubmit}>
    <p>ワークスペースの準備ができました！✨</p>
    <button
      type="submit"
      className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
    >
      ワークスペースを開始する
    </button>
    {error && <div className="text-red-600 text-sm">{error}</div>}
  </form>
);
