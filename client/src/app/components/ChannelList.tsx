import Link from "next/link";

type Channel = {
  id: number;
  name: string;
};

type ChannelListProps = {
  channels: Channel[];
};

export default function ChannelList({ channels }: ChannelListProps) {
  return (
    <ul>
      {channels &&
        channels.map((channel) => (
          <li key={channel.id}>
            <Link
              href={`/channels/${channel.id}`}
              className="block hover:bg-gray-200 p-2 rounded"
            >
              # {channel.name}
            </Link>
          </li>
        ))}
    </ul>
  );
}
