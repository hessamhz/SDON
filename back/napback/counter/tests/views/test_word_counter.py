from counter.forms import WordCountForm
from django.test import Client, TestCase
from django.urls import reverse
from faker import Faker


class WordCountViewTest(TestCase):
    """
    Test case for the word count view.
    """

    def setUp(self):
        """
        Set up the necessary objects and configurations for the tests.
        """
        self.client = Client()
        self.faker = Faker()

    def test_get_request(self):
        """
        Test the GET request to the word count view.
        The view should return a response with status code 200, use the correct template,
        and have a WordCountForm instance in the response context.
        """
        response = self.client.get(reverse("word_count"))
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, "counter.html")
        self.assertIsInstance(response.context.get("form"), WordCountForm)

    def test_post_request_with_valid_form(self):
        """
        Test the POST request to the word count view with a valid form submission.
        The view should return a response with status code 200, use the correct template,
        and have the correct word count in the response context.
        """
        fake_text = self.faker.paragraph()
        form_data = {"text": fake_text}
        response = self.client.post(reverse("word_count"), data=form_data)
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, "counter.html")
        self.assertEqual(
            response.context.get("word_count"), len(fake_text.split())
        )

    def test_post_request_with_empty_form(self):
        """
        Test the POST request to the word count view with an empty form submission.
        The view should return a response with status code 200, use the correct template,
        have a WordCountForm instance in the response context, and display form errors.
        """
        form_data = {"text": ""}
        response = self.client.post(reverse("word_count"), data=form_data)
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, "counter.html")
        self.assertIsInstance(response.context.get("form"), WordCountForm)
        self.assertTrue(response.context.get("form").errors)
