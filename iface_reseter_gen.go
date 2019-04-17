// Code generated by generate/interfaces; DO NOT EDIT.

package disgord

func (c *Channel) Reset() {
	c.ID = 0
	c.Type = 0
	c.GuildID = 0
	c.Position = 0
	c.PermissionOverwrites = nil
	c.Name = ""
	c.Topic = ""
	c.NSFW = false
	c.LastMessageID = 0
	c.Bitrate = 0
	c.UserLimit = 0
	c.RateLimitPerUser = 0
	c.Recipients = nil
	c.Icon = ""
	c.OwnerID = 0
	c.ApplicationID = 0
	c.ParentID = 0
	c.LastPinTimestamp = Time{}
	c.complete = false
	c.recipientsIDs = nil
}

func (e *Emoji) Reset() {
	e.mu = Lockable{}
	e.ID = 0
	e.Name = ""
	e.Roles = nil
	if e.User != nil {
		e.User.Reset()
	}
	e.RequireColons = false
	e.Managed = false
	e.Animated = false
	e.guildID = 0
}

func (m *MessageCreate) Reset() {
	if m.Message != nil {
		m.Message.Reset()
	}
	m.Ctx = nil
	m.ShardID = 0
}

func (g *Guild) Reset() {
	g.ID = 0
	g.ApplicationID = 0
	g.Name = ""
	g.Icon = nil
	g.Splash = nil
	g.Owner = false
	g.OwnerID = 0
	g.Permissions = 0
	g.Region = ""
	g.AfkChannelID = 0
	g.AfkTimeout = 0
	g.EmbedEnabled = false
	g.EmbedChannelID = 0
	g.VerificationLevel = 0
	g.DefaultMessageNotifications = 0
	g.ExplicitContentFilter = 0
	g.Roles = nil
	g.Emojis = nil
	g.Features = nil
	g.MFALevel = 0
	g.WidgetEnabled = false
	g.WidgetChannelID = 0
	g.SystemChannelID = 0
	g.JoinedAt = nil
	g.Large = false
	g.Unavailable = false
	g.MemberCount = 0
	g.VoiceStates = nil
	g.Members = nil
	g.Channels = nil
	g.Presences = nil
}

func (m *Member) Reset() {
	m.GuildID = 0
	if m.User != nil {
		m.User.Reset()
	}
	m.Nick = ""
	m.Roles = nil
	m.JoinedAt = Time{}
	m.Deaf = false
	m.Mute = false
	m.userID = 0
}

func (m *Message) Reset() {
	m.ID = 0
	m.ChannelID = 0
	if m.Author != nil {
		m.Author.Reset()
	}
	m.Content = ""
	m.Timestamp = Time{}
	m.EditedTimestamp = Time{}
	m.Tts = false
	m.MentionEveryone = false
	m.Mentions = nil
	m.MentionRoles = nil
	m.Attachments = nil
	m.Embeds = nil
	m.Reactions = nil
	m.Nonce = 0
	m.Pinned = false
	m.WebhookID = 0
	m.Type = 0
	m.Activity = MessageActivity{}
	m.Application = MessageApplication{}
	m.GuildID = 0
	m.SpoilerTagContent = false
	m.SpoilerTagAllAttachments = false
	m.HasSpoilerImage = false
}

func (r *Reaction) Reset() {
	r.Count = 0
	r.Me = false
	r.Emoji = nil
}

func (r *Role) Reset() {
	r.ID = 0
	r.Name = ""
	r.Color = 0
	r.Hoist = false
	r.Position = 0
	r.Permissions = 0
	r.Managed = false
	r.Mentionable = false
	r.guildID = 0
}

func (a *Activity) Reset() {
	a.Name = ""
	a.Type = 0
	a.URL = nil
	a.Timestamps = nil
	a.ApplicationID = 0
	a.Details = nil
	a.State = nil
	a.Party = nil
	a.Assets = nil
	a.Secrets = nil
	a.Instance = false
	a.Flags = 0
}

func (u *User) Reset() {
	u.ID = 0
	u.Username = ""
	u.Discriminator = 0
	u.Email = ""
	u.Avatar = nil
	u.Token = ""
	u.Verified = false
	u.MFAEnabled = false
	u.Bot = false
	u.PremiumType = 0
}

func (v *VoiceState) Reset() {
	v.GuildID = 0
	v.ChannelID = 0
	v.UserID = 0
	if v.Member != nil {
		v.Member.Reset()
	}
	v.SessionID = ""
	v.Deaf = false
	v.Mute = false
	v.SelfDeaf = false
	v.SelfMute = false
	v.Suppress = false
}

func (v *VoiceRegion) Reset() {
	v.ID = ""
	v.Name = ""
	v.SampleHostname = ""
	v.SamplePort = 0
	v.VIP = false
	v.Optimal = false
	v.Deprecated = false
	v.Custom = false
}
