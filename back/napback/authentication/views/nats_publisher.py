import asyncio
from nats.aio.client import Client as NATS
from django.conf import settings

async def publish_message(subject, message):
    nc = NATS()
    await nc.connect(servers=[settings.NATS_URL])  # Replace with your NATS server address
    await nc.publish(subject, message.encode('utf-8'))
    await nc.flush()
    await nc.close()

def send_nats_message(subject, message):
    asyncio.run(publish_message(subject, message))
    print("message sended")
    print(subject,message)
