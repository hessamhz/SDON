import asyncio
from nats.aio.client import Client as NATS

async def run():
    nc = NATS()

    await nc.connect(servers=["nats://localhost:4222"])

    async def message_handler(msg):
        subject = msg.subject
        reply = msg.reply
        data = msg.data.decode()
        print(f"Received a message on '{subject} {reply}': {data}")

    # Subscribe to the subjects you expect Django to publish to
    await nc.subscribe("create.infrastructure", cb=message_handler)
    await nc.subscribe("create.service", cb=message_handler)
    await nc.subscribe("delete", cb=message_handler)
    await nc.subscribe("delete", cb=message_handler)
    await nc.subscribe("delete", cb=message_handler)
    await nc.subscribe("visualize.service",cb=message_handler)
    await nc.subscribe("visualize.infrastructure",cb=message_handler)
    # Add more subscriptions as necessary

    print("Listening for messages on subscribed subjects...")

    while True:
        await asyncio.sleep(10)  # Sleeps to keep the loop running

if __name__ == '__main__':
    asyncio.run(run())
