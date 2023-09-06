"use client";
import Image from "next/image";
import wficon from "../icon.svg";
import { pb } from "../../lib/pocketbase";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function Home() {
  const { push } = useRouter();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const usernameChange = (e: any) => {
    setUsername(e.target.value);
  };

  const passwordChange = (e: any) => {
    setPassword(e.target.value);
  };

  const loginSubmit = async (e: any) => {
    e.preventDefault();
    const authData = await pb
      .collection("wf_users")
      .authWithPassword(username, password);
    document.cookie = pb.authStore.exportToCookie({ httpOnly: false });

    if (authData) {
      push("/");
    }

    if (pb.authStore.isValid) {
    } else {
    }
  };

  return (
    <main className="min-h-screen flex flex-col justify-center items-center">
      <Image
        priority
        src={wficon}
        alt="Wiredfriends"
        width={64}
        className="m-3 pointer-events-none"
      />

      <form
        onSubmit={loginSubmit}
        className="opacity-100 transition-all duration-1000 rounded-xl transition flex flex-col"
      >
        <input
          type="text"
          placeholder="User Name"
          onChange={usernameChange}
          className="p-2 m-1 rounded dark:bg-slate-600"
        />
        <input
          type="password"
          placeholder="Pass Word"
          onChange={passwordChange}
          className="p-2 m-1 rounded dark:bg-slate-600"
        />

        <button
          type="submit"
          className="flex-1 p-3 m-2 rounded-xl bg-slate-800 text-blue-100 hover:text-blue-300 hover:bg-slate-700 dark:bg-slate-300 dark:text-slate-700 dark:hover:bg-slate-200 dark:hover:text-slate-700"
        >
          Log In
        </button>
      </form>
    </main>
  );
}
