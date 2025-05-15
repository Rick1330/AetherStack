export default function HomePage() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24 bg-gradient-to-br from-slate-900 to-slate-800">
      <div className="text-center">
        <h1 className="text-6xl font-bold text-sky-400 drop-shadow-lg">
          AetherStack Dashboard
        </h1>
        <p className="mt-4 text-2xl text-slate-300">
          Dashboard Online
        </p>
        <div className="mt-8 p-6 bg-slate-700/50 rounded-lg shadow-xl">
          <p className="text-slate-400">
            {/* TODO: Replace with actual dashboard content */}
            Welcome to the AetherStack dashboard. This area will soon display repository status, AI agent reports, and AI-generated insights.
          </p>
        </div>
      </div>
    </main>
  );
}

