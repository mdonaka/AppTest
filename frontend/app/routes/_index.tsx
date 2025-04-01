import type {MetaFunction} from "@remix-run/node";
import "~/styles/Index.scss";

export const meta: MetaFunction = () => {
  return [
    {title: "Challenge the Quiz!"},
    {name: "description", content: "Challenge the Quiz!"},
  ];
};

export default function Index() {
  return (
    <div className="container">
      <h1 className="title">Challenge the Quiz!</h1>
    </div>
  );
}
