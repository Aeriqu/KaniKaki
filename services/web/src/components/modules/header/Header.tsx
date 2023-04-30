import Image from 'next/image';
import Link from 'next/link';
import kanikakiBanner from '../../../assets/images/kanikaki-banner.png';
import HeaderLoginLogout from '@/components/elements/header-login-logout/HeaderLoginLogout';
import ClientWrapper from '@/components/elements/client-wrapper/ClientWrapper';

export default function Header() {
  return (
    <header className=''>
      <nav className='flex items-center justify-between mx-auto max-w-7xl p-6'>
        <div className='flex justify-start lg:flex-1'>
          <Link href='/'>
            <Image className='h-8 w-auto' src={kanikakiBanner} alt='KaniKaki Home' />
          </Link>
        </div>

        <div className='lg:flex lg:flex-1 lg:justify-end'>
          <ClientWrapper><HeaderLoginLogout></HeaderLoginLogout></ClientWrapper>
        </div>
      </nav>
    </header>
  )
}