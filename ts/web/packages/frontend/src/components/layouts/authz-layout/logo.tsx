import Image from 'next/image'

const Logo = () => {
  return (
    <div className="flex items-center">
      <Image
        src="/images/logo.jpg"
        alt="ロゴ"
        width={32}
        height={32}
        className="mr-2"
      />
      <span className="text-xl font-bold text-gray-900">Stack Example</span>
    </div>
  )
}

export { Logo }
