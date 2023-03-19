interface Message {
  sender: string;
  sentOn: string;
  message: string;
}

interface MapRenderableMessage extends Message {
  id: number;
}

export type { Message, MapRenderableMessage };
