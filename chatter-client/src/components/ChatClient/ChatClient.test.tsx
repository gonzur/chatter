import { fireEvent, render, screen } from "@testing-library/react";
import { keyboard } from "@testing-library/user-event/dist/keyboard";

import ChatCLient from "./ChatCLient";

test("accepts a message", () => {
  render(<ChatCLient userID="test" />);

  const enter = screen.getByRole("button");
  const inputBox = screen.getByRole("textbox");

  fireEvent.click(inputBox);
  keyboard("hello");
  fireEvent.click(enter);

  expect(screen).toHaveTextContent();
});
