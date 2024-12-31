/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 * SPDX-License-Identifier: MIT-0
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this
 * software and associated documentation files (the "Software"), to deal in the Software
 * without restriction, including without limitation the rights to use, copy, modify,
 * merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
 * INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
 * PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
 * HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 * OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package com.amazon.sample.orders.messaging.sqs;

import com.amazon.sample.orders.messaging.MessagingProvider;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import io.awspring.cloud.sqs.operations.SqsOperations;

public class SqsMessagingProvider implements MessagingProvider {

    private final String messageQueueTopic;
    private final SqsOperations blockingTemplate;
    private final ObjectMapper mapper;

    public SqsMessagingProvider(String messageQueueTopic, SqsOperations blockingTemplate, ObjectMapper mapper) {
        this.blockingTemplate = blockingTemplate;
        this.messageQueueTopic = messageQueueTopic;
        this.mapper = mapper;
    }

    @Override
    public void publishEvent(Object event) {
        blockingTemplate.send(to -> {
            try {
                to.queue(messageQueueTopic)
                    .payload(mapper.writeValueAsString(event))
                    .delaySeconds(10);
            } catch (JsonProcessingException e) {
                e.printStackTrace();
            }
        });
    }
}