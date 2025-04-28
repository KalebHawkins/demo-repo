def lambda_handler(event, context):
    print("Notifying admins...")
    return { "message": "Admins notified!" }
