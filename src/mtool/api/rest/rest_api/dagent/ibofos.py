'''
/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Intel Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
 '''
 
import os
import requests
import uuid
import time
import json
from flask import make_response
import logging
from logging.handlers import RotatingFileHandler

ip = os.environ.get('DAGENT_HOST', 'localhost')
port = '3000'

array_names= [ 'POSArray', 'POSArray2' ]
DAGENT_URL = 'http://' + ip + ':' + port
BASE_PATH = 'api/ibofos'
BASIC_AUTH_TOKEN = 'Basic YWRtaW46YWRtaW4='
VERSION = 'v1'

connect_timeout = 60
read_timeout = 60


def get_headers(
        auth=BASIC_AUTH_TOKEN,
        content_type="application/json"):
    return {"X-Request-Id": str(uuid.uuid4()),
            "Accept": content_type,
            "Authorization": auth,
            "ts": str(int(time.time()))}


def make_failure_response(desc='unable to perform the task', code=500):
    return make_response(json.dumps({"error": desc}), code)


def send_command_to_dagent(req_type, url, headers, timeout=None, data=None):
    retry_count = 0
    response = None
    try:
        while(1):
            #print("sending request to dagent ",url)
            if(req_type == "GET"):
                response = requests.get(
                    url=url, headers=headers, timeout=timeout)
            elif(req_type == "POST"):
                response = requests.post(
                    url=url, headers=headers, timeout=timeout, data=data)
            elif(req_type == "DELETE"):
                response = requests.delete(
                    url=url, headers=headers, timeout=timeout, data=data)
            elif(req_type == "PUT"):
                response = requests.put(
                    url=url, headers=headers, timeout=timeout, data=data)
            elif(req_type == "PATCH"):
                response = requests.patch(
                    url=url, headers=headers, timeout=timeout, data=data)
            retry_count = retry_count + 1
            if((response.status_code == 200 and response['result'] and response['result']['status'] and response['result']['status']['code'] and response['result']['status']['code'] != '12000')):
                return response
            if response.status_code == 202 or response.status_code == 503:
                retry_count = 5
            if(retry_count >= 5):
                return response
            time.sleep(1)
    except BaseException:
        #print("in except block", url)
        return response


def get_system_state(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Get system state...')
    try:
        response = send_command_to_dagent("GET",
                                          url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'system',
                                          headers={"X-Request-Id": str(uuid.uuid4()),
                                                   "Accept": "application/json",
                                                   "Authorization": auth,
                                                   "ts": str(int(time.time()))},
                                          timeout=(connect_timeout,
                                                   read_timeout))
        #print("response:",response.status_code, response.json())
        # if "error" in response.json():
        #    logger.error('%s %s', 'ERROR', response)
        #    return {"result": response["error"], "return": -1}
        # else:
        #    logger.info('%s %s', 'INFO sysstate', json.dumps(response))
        return response
    except Exception as err:
        print(f'Other error occurred in get_system_state : {err}', err)
    return {"result": "could not get the system state", "return": -1}

"""
def get_dagent_state(auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL +
            '/api/dagent/v1/ping',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("--------------RESPONSE-------------")
        response = response.json()
        #print(response['result']['status']['code'])
        #print("--------------RESPONSE-------------")
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
"""

def start_ibofos(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to start ibofos...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'system',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        #array_exists(array_names[0])
        return response
    except Exception as err:
        print(f'Other error occurred start_ibofos: {err}')

    return make_failure_response('Could not get ibofos to start...', 500)


def stop_ibofos(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to stop ibofos...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'system',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    # finally:
        # return 'Could not get ibofos to stop... 500'
    return make_failure_response('Could not get ibofos to stop...', 500)

"""
def exit_system(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to exit ibofos...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL +
            '/api/ibofos/v1/system/exitibofos',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    # finally:
        # return 'Could not get ibofos to stop... 500'
    return make_failure_response('Could not get ibofos to exit...', 500)
"""

def scan_devices(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to scan devices...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'devices/all/scan',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------",response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to scan devices...', 500)


def get_devices(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to list devices ...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'devices',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE get devices---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to scan devices...', 500)

def get_smart_info(name, auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to list devices ...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'devices/' + name + '/smart',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to scan devices...', 500)


def delete_array(name, auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to delete array...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
                "DELETE",
                url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'array/' + name,
                headers=req_headers,
                timeout=(
                    connect_timeout,
                    read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to get array status...', 500)


def array_status(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to get array status...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'system',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to get array status...', 500)


def array_exists(arrayname=array_names[0],auth=BASIC_AUTH_TOKEN):
    #print("in array exists")
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to get array status...')
    req_headers = get_headers(auth)
    try:
        scan_devices()

        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'array/' + arrayname + '/load',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        
    
        if (response.status_code == 200):
            response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'system/mount',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout)
            )
            return True
        else:
            return False

    except Exception as e:
        print("exception in exists: " + str(e))
        return False


def create_array(
        name,
        raidtype,
        spare_devices,
        data_devices,
        meta_devices,
        auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info(
        '%s',
        'Sending command to D-Agent to create array using mount api...')
    req_headers = get_headers(auth)
    request_body = {
        "param": {
            "name": name,
            "raidtype": "RAID5",
            "buffer": meta_devices,
            "data": data_devices,
            "spare": spare_devices}}
    request_body = json.dumps(request_body)
    #print(request_body)
    try:
        #print("request body create array ", request_body)
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'array',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("resp for create array",response.json())
        if response.status_code != 200:
            return response
        response = mount_array(name)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to scan devices...', 500)
def mount_array(arrayname, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {
            "param": {
                "array": arrayname
            }
    }
    request_body = json.dumps(request_body)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'array/' + arrayname + '/mount',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code , response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not mount array.', 500)
def unmount_array(arrayname, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {
            "param": {
                "array": arrayname
            }
    }
    request_body = json.dumps(request_body)
    try:
        response = send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'array/' + arrayname + '/mount',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not unmount array', 500)

def qos_create_volume_policies(request_body, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = json.dumps(request_body)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/qos',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not create qos policies', 500)
def qos_reset_volume_policies(request_body, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = json.dumps(request_body)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/qos/reset',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not reset qos policies', 500)

def qos_list_volume_policies(request_body, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = json.dumps(request_body)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/qos/policies',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get qos policies', 500)


def array_info(array_name, auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to get list array...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'array/' + array_name,
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------",response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to get list array...', 500)
def list_arrays(auth=BASIC_AUTH_TOKEN):
    logger = logging.getLogger(__name__)
    logger.info('%s', 'Sending command to D-Agent to get list arrays...')
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'arrays',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------",response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to get list array...', 500)


def create_volume(
        name,
        arrayname,
        size,
        count,
        suffix,
        mount_all,
        stop_on_error,
        maxbw=0,
        maxiops=0,
        auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {
        "param": {
            "name": name,
            "array": arrayname,
            "size": size,
            "maxbw": maxbw,
            "maxiops": maxiops,
            "totalcount": count,
            "stoponerror": stop_on_error,
            "namesuffix": suffix,
            "mountall": mount_all}}

    request_body = json.dumps(request_body)
    #print("volume bodyyyy",request_body)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------",response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to scan devices...', 500)


def update_volume(params, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {"param": params}
    request_body = json.dumps(request_body)
    #print(request_body)
    try:
        response = send_command_to_dagent(
            "PATCH",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes/' + params["name"] + '/qos',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to scan devices...', 500)


def rename_volume(params, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = json.dumps(params)
    try:
        response = send_command_to_dagent(
            "PATCH",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes/' + params["param"]["name"],
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RENAME RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Error response from Agent while renaming volume...', 500)


def mount_volume(name, arrayname, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {
            "param": {
                "array": arrayname
            }
    }
    request_body = json.dumps(request_body)
    #print(request_body)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes/' + name + '/mount',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code , response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to scan devices...', 500)


def unmount_volume(name, arrayname, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {
            "param": {
                "array": arrayname
            }
    }
    request_body = json.dumps(request_body)
    #print(request_body)
    try:
        response = send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes/' + name + '/mount',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to unmount volumes...', 500)

def mount_ibofos(auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)

    try:
        method = 'system/mount'
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + method,
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))

        #print("---------------RESPONSE---------------")
        #print(response.status_code , response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not mount POS...', 500)

def unmount_ibofos(auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)

    try:
        method = 'system/mount'
        response = send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + method,
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))

        #print("---------------RESPONSE---------------",response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not unmount POS...', 500)


def list_volumes(array_name, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/'+ 'volumelist/'+array_name,
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code , response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to list volumes...', 500)


def delete_volume(name, arrayname, auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {
            "param": {
                "name": name,
                "array": arrayname
            }
    }
    request_body = json.dumps(request_body)
    
    try:
        send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes/' + name + '/mount',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)

        response = send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes/' + name,
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get ibofos to delete volumes...', 500)


def max_vol_count(auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "GET",
            url=DAGENT_URL + '/' + BASE_PATH + '/' + VERSION + '/' + 'volumes/maxcount',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get max vol count from POS...', 500)





def add_spare_disk(name, arrayname=array_names[0], auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    request_body = {
            "param": {
                "spare":[
                    {
                        "deviceName" : name,
                    }
                 ]
            }
    }
    request_body = json.dumps(request_body)
    #print(request_body)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL + '/' +
            BASE_PATH + '/' + VERSION + '/array/' + arrayname + '/devices',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout),
            data=request_body)
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get POS to add a spare disk...', 500)


def remove_spare_disk(name, arrayname=array_names[0], auth=BASIC_AUTH_TOKEN):
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "DELETE",
            url=DAGENT_URL + '/' + 
            BASE_PATH + '/' + VERSION + '/array/' + arrayname + '/devices/' + name,
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
        return response
    except Exception as err:
        print(f'Other error occurred: {err}')
    return make_failure_response(
        'Could not get POS to remove spare disk...', 500)

"""
def report_test(auth=BASIC_AUTH_TOKEN):
    # •Report file directory : /etc/ibofos/report
    req_headers = get_headers(auth)
    try:
        response = send_command_to_dagent(
            "POST",
            url=DAGENT_URL +
            '/api/ibofos/v1/test/report',
            headers=req_headers,
            timeout=(
                connect_timeout,
                read_timeout))
        #print("---------------RESPONSE---------------")
        #print(response.status_code, response.json())
    except Exception as err:
        print(f'Other error occurred: {err}')
"""

handler = RotatingFileHandler(
    'log/ibof.log',
    maxBytes=1024 * 1024,
    backupCount=3)
log_handler = logging.getLogger(__name__)
log_handler.setLevel(logging.INFO)
log_handler.addHandler(handler)