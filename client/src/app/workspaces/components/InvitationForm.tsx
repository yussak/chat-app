import { FormEvent } from "react";

type InvitationFormProps = {
  error: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => Promise<void>;
};

export const InvitationForm = ({ error, onSubmit }: InvitationFormProps) => (
  <form className="mt-8 space-y-6" onSubmit={onSubmit}>
    <p>手順 3/5</p>
    <div>
      <label
        htmlFor="invitation"
        className="block text-sm font-medium text-gray-700"
      >
        一緒に仕事をする人をメールアドレスで追加する
      </label>
      <textarea
        id="invitation"
        name="invitation"
        rows={3}
        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        placeholder="例: ellis@gmail.com、 maria@gmail.com"
      />
    </div>

    {error && <div className="text-red-600 text-sm">{error}</div>}

    <div className="flex gap-4">
      <button className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
        次へ（未実装）
      </button>
      <button
        type="submit"
        className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        この手順をスキップする
      </button>
    </div>
  </form>
);
