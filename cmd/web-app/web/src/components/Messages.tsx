import { formatISO } from "date-fns";
import type { Message } from "../hooks/use-drone";

interface MessagesProps {
  messages?: Array<Message>;
}

const MessageRow = (props: { message: Message }) => {
  const { message } = props;
  return (
    <p>
      <strong>
        {formatISO(new Date(message.time), {
          representation: "time",
          format: "extended",
        })}
      </strong>{" "}
      {message.message}
    </p>
  );
};

export const Messages = (props: MessagesProps) => {
  const { messages } = props;

  return (
    <div className="px-6 py-5 liquid-bg flex flex-col rounded-4xl grow mt-4">
      <span className="font-black text-xl">Messages</span>
      {!messages?.length ? <p>No messages</p> : null}
      {messages?.length
        ? messages
            .slice(Math.max(messages.length - 4, 0))
            .map((message) => (
              <MessageRow key={message.time} message={message} />
            ))
        : null}
    </div>
  );
};
