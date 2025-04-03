import Link from "next/link";
import { usePathname } from "next/navigation";

type Channel = {
  id: number;
  workspace_id: number;
  name: string;
  is_public: boolean;
};

type ChannelListProps = {
  channels: Channel[];
  workspaceId: number;
};

export default function ChannelList({
  channels,
  workspaceId,
}: ChannelListProps) {
  const pathname = usePathname();
  const currentChannelId = pathname.split("/").pop();

  return (
    <ul>
      {channels &&
        channels.map((channel) => (
          <li key={channel.id}>
            <Link
              href={`/workspaces/${workspaceId}/channels/${channel.id}`}
              className={`block hover:bg-gray-200 p-2 rounded ${
                currentChannelId === channel.id.toString() ? "text-red-500" : ""
              }`}
            >
              # {channel.name}
            </Link>
          </li>
        ))}
    </ul>
  );
}
