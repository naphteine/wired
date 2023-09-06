import Image from "next/image";
import wficon from "./icon.svg";
import { pb } from "../lib/pocketbase";
import Link from "next/link";

export default function Home() {
  const user = pb.authStore.model;
  //const cookie = req.cookies.get("pb_auth");
  //pb.authStore.loadFromCookie(cookie || "");
  //console.log(cookie);
  console.log(" --------------------------- \n", pb.authStore);

  return (
    <main className="min-h-screen flex flex-col justify-center items-center">
      <Image
        priority
        src={wficon}
        alt="Wiredfriends"
        width={64}
        className="m-3 pointer-events-none"
      />

      {user && user.email}

      {pb.authStore.isValid ? (
        <h1>Hello {pb.authStore.model?.email}</h1>
      ) : (
        <Link href="/auth">Authorize Link</Link>
      )}
    </main>
  );
}
