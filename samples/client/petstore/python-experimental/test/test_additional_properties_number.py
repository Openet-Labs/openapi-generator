# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    OpenAPI spec version: 1.0.0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import petstore_api
from petstore_api.models.additional_properties_number import AdditionalPropertiesNumber  # noqa: E501
from petstore_api.rest import ApiException
from petstore_api.exceptions import ApiTypeError


class TestAdditionalPropertiesNumber(unittest.TestCase):
    """AdditionalPropertiesNumber unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAdditionalPropertiesNumber(self):
        """Test AdditionalPropertiesNumber"""
        # can make model without additional properties
        model = AdditionalPropertiesNumber()

        # can make one with additional properties
        model = AdditionalPropertiesNumber(some_key=11.3)
        assert model['some_key'] == 11.3

        # type checking works on additional properties
        with self.assertRaises(ApiTypeError) as exc:
            model = AdditionalPropertiesNumber(some_key=10)

if __name__ == '__main__':
    unittest.main()
