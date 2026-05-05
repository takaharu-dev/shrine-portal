import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Jinja Journey",
  description: "Shrine portal development frontend",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}
