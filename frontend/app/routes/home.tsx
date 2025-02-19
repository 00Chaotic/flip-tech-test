import type { Route } from "./+types/home";
import { Shop } from "~/shop/shop";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Flip Tech Test" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Home() {
  return <Shop />;
}
