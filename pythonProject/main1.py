





import logging

import boto3
from botocore.exceptions import ClientError

logger = logging.getLogger(__name__)
# sqs = boto3.resource('sqs')
session = boto3.Session()
sqs = session.resource('sqs')


def get_queues(prefix=None):
    if prefix:
        queue_iter = sqs.queues.filter(QueueNamePrefix=prefix)
    else:
        queue_iter = sqs.queues.all()
    queues = list(queue_iter)
    if queues:
        logger.info("Got queues: %s", ', '.join([q.url for q in queues]))
    else:
        logger.warning("No queues found.")
    return queues


if __name__ == '__main__':
    get_queues()
