"use client";

import { useParams } from "next/navigation";
import { useState } from "react";
import { useEffect } from "react";
import { api } from "@/app/lib/api-client";

export default function Channel() {
  const params = useParams();
  const id = params.id;

  const [channel, setChannel] = useState(null);

  useEffect(() => {
    const fetchChannel = async () => {
      const res = await api.get(`/channels/${id}`);
      setChannel(res.data);
    };
    fetchChannel();
  }, [id]);

  return (
    <div>
      Channel {id}
      <br />
      channel name: {channel?.name}
    </div>
  );
}
