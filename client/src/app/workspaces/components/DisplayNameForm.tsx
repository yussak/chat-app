import { FormEvent } from "react";

type DisplayNameFormProps = {
  displayName: string;
  setDisplayName: (email: string) => void;
  error: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => Promise<void>;
};

export const DisplayNameForm = ({
  displayName,
  setDisplayName,
  error,
  onSubmit,
}: DisplayNameFormProps) => (
  <form className="mt-8 space-y-6" onSubmit={onSubmit}>
    <p>手順 2/5</p>
    <div>
      <label
        htmlFor="displayName"
        className="block text-sm font-medium text-gray-700"
      >
        表示名
      </label>
      <input
        id="displayName"
        name="displayName"
        type="text"
        required
        value={displayName}
        onChange={(e) => setDisplayName(e.target.value)}
        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        placeholder="例: 山田 太郎"
      />
    </div>

    {error && <div className="text-red-600 text-sm">{error}</div>}

    <div className="flex gap-4">
      <button
        type="submit"
        className="flex-1 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        次へ
      </button>
    </div>
  </form>
);
