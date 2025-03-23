import { FormEvent } from "react";

type ThemeFormProps = {
  theme: string;
  setTheme: (theme: string) => void;
  error: string;
  onSubmit: (e: FormEvent<HTMLFormElement>) => Promise<void>;
};

export const ThemeForm = ({
  theme,
  setTheme,
  error,
  onSubmit,
}: ThemeFormProps) => (
  <form className="mt-8 space-y-6" onSubmit={onSubmit}>
    <p>手順 4/5</p>
    <div>
      <p className="block text-sm font-medium text-gray-700">
        チームで取り組んでいることを教えてください。プロジェクト、キャンペーン、イベント、まとめようとしている案件など、いろいろなことが考えられます。
      </p>
      <input
        id="theme"
        name="theme"
        type="text"
        required
        value={theme}
        placeholder="例: 第４半期予算、秋のキャンペーン"
        onChange={(e) => setTheme(e.target.value)}
        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
      />
      {error && <div className="text-red-600 text-sm">{error}</div>}

      <button
        type="submit"
        className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
      >
        次へ
      </button>
    </div>
  </form>
);
