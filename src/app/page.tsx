import Image from "next/image";
import wficon from "./icon.svg";

export default function Home() {
  return (
    <main className="min-h-screen flex justify-center items-center">
      <Image priority src={wficon} alt="Wiredfriends" width={64} />
    </main>
  );
}
