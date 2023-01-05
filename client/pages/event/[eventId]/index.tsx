import Head from 'next/head';
import { Inter } from '@next/font/google';
import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../components/templates/shared/MyHead';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { EventInfo } from '../../../components/molecules/EventInfo';

const inter = Inter({ subsets: ['latin'] });

export default function Show() {
  const [isHost, setIsHost] = useState(false);
  const [hasJoined, setHasJoined] = useState(false);
  const event_id = useRouter().query.id;
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <BasicTemplate className="text-center">
        <EventInfo />
      </BasicTemplate>
    </>
  );
}
