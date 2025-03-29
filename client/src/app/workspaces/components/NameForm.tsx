import { FormEvent } from "react";

type NameFormProps = {
  name: string;
  setName: (name: string) => void;
  error: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => Promise<void>;
};

export const NameForm = ({ name, setName, error, onSubmit }: NameFormProps) => (
  <form className="mt-8 space-y-6" onSubmit={onSubmit}>
    <p>手順 1/5</p>
    <div>
      <input
        id="name"
        name="name"
        type="text"
        required
        value={name}
        onChange={(e) => setName(e.target.value)}
        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
        placeholder="例: ABC 営業部、ABC 社"
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
