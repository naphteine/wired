"use client";

import Image from "next/image";
import wficon from "./icon.svg";
import { useState } from "react";

export default function Home() {
  const [auth, setAuth] = useState(false);

  function authorize(e: any) {
    setAuth(true);
  }

  return (
    <main className="min-h-screen flex flex-col justify-center items-center">
      <Image
        priority
        src={wficon}
        alt="Wiredfriends"
        width={64}
        className="m-3 pointer-events-none"
      />

      <div
        className={`opacity-100 transition-all duration-1000 rounded-xl transition flex flex-col ${
          !auth && "opacity-0 hidden"
        }`}
      >
        <input
          type="text"
          placeholder="User Name"
          className="p-2 m-1 rounded"
        />
        <input
          type="password"
          placeholder="Pass Word"
          className="p-2 m-1 rounded"
        />

        <div className="flex">
          <button className="flex-1 p-3 m-2 rounded-xl bg-slate-800 text-blue-100 hover:text-blue-300 hover:bg-slate-700">
            Sign Up
          </button>

          <button className="flex-1 p-3 m-2 rounded-xl bg-slate-800 text-blue-100 hover:text-blue-300 hover:bg-slate-700">
            Log In
          </button>
        </div>
      </div>

      <button
        onClick={authorize}
        className={`transition-all opacity-100 duration-1000 p-3 m-2 rounded-xl bg-slate-800 text-blue-100 hover:text-blue-300 hover:bg-slate-700 ${
          auth && "opacity-0 hidden"
        }`}
      >
        Authorize
      </button>
    </main>
  );
}
