import { cleanup, fireEvent, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import ChatCLient from "./ChatCLient";

test("accepts a message", () => {
  render(<ChatCLient username="test" roomID="test" />);

  const messages = [
    "hello",
    "what a fine evening this is",
    "the quick brown fox jumps over the lazy dog",
    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    "and so on.",
  ];

  const enter = screen.getByRole("button");
  const inputBox = screen.getByRole("textbox");

  messages.forEach((message) => {
    userEvent.type(inputBox, message);
    fireEvent.click(enter);
  });

  const renderedMessages = screen.getAllByTestId("chatMessage");

  messages.forEach((message, index) => {
    expect(renderedMessages.at(index)).toHaveTextContent(message);
  });
});

test("displays user next to message", () => {
  const userIDs = [
    "jameson",
    "asdf asdfsdf",
    "werd alllooo cna",
    "..............................",
  ];

  userIDs.forEach((id) => {
    render(<ChatCLient username={id} roomID={id} />);

    const inputBox = screen.getByRole("textbox");
    const enter = screen.getByRole("button");

    userEvent.type(inputBox, id);
    fireEvent.click(enter);

    expect(screen.getByTestId("chatUser")).toHaveTextContent(id);
    cleanup();
  });
});
