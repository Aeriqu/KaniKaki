import { ErrorBannerProp } from "@/types/ErrorBannerProp";

export function ErrorBanner({ error }: Readonly<ErrorBannerProp>) {
  if (!error) {
    return (
      <></>
    )
  }

  return (
    <div className="bg-red-500 rounded-lg w-full p-4">
      {error}
    </div>
  )
}