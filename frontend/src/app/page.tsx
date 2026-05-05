const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";

export default function Home() {
  return (
    <main className="min-h-screen bg-shrine-paper text-shrine-ink">
      <section className="mx-auto flex min-h-screen max-w-4xl flex-col justify-center px-6 py-16">
        <p className="text-sm font-semibold uppercase tracking-[0.18em] text-shrine-green">
          Jinja Journey
        </p>
        <h1 className="mt-4 text-4xl font-bold leading-tight sm:text-5xl">
          神社巡りをもっと自由に。
        </h1>
        <p className="mt-6 max-w-2xl text-base leading-8 text-neutral-700">
          Next.js、Go、PostgreSQL を Docker Compose で起動するための最小構成です。
          API のヘルスチェックは backend の <code className="rounded bg-white px-1 py-0.5">/health</code> で確認できます。
        </p>
        <div className="mt-8 flex flex-col gap-3 text-sm sm:flex-row">
          <a
            className="inline-flex items-center justify-center rounded-md bg-shrine-red px-4 py-3 font-semibold text-white"
            href={`${apiBaseUrl}/health`}
          >
            Backend health
          </a>
          <div className="inline-flex items-center rounded-md border border-neutral-200 bg-white px-4 py-3 text-neutral-700">
            DB: localhost:5432
          </div>
        </div>
      </section>
    </main>
  );
}
