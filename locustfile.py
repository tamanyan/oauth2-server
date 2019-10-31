# -*- coding: utf-8 -*-
from __future__ import absolute_import
from __future__ import unicode_literals

from locust import HttpLocust, TaskSet, task


class UserTaskSet(TaskSet):
    def on_start(self):
        """
        タスクセットの開始時に1回のみ呼ばれます。
        """
        # self.client.post("/login", {"username": "ellen_key", "password": "education"})
        pass

    @task
    def token(self):
        self.client.get("/oauth2/token")

    # @task
    # def index(self):
    #     self.client.get("/")


class WebsiteUser(HttpLocust):
    task_set = UserTaskSet

    min_wait = 10000
    max_wait = 10000

