import type {MetaFunction} from "@remix-run/node";
import {Link} from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    {title: "New Remix App"},
    {name: "description", content: "Welcome to Remix!"},
  ];
};


export default function Index() {
  return (
    <div>
      <Link to="/list">Go to About</Link>
    </div>
  )
}
